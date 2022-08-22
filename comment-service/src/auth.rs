use jwt::{Token, Header, Unverified};

use rocket::Outcome;
use rocket::http::Status;
use rocket::request::{self, Request, FromRequest};

use reqwest;

#[derive(Serialize, Deserialize)]
pub struct MyJWTClaims {
    pub email: String,
    pub role: String,
    pub exp: i64,
}

impl<'a, 'r> FromRequest<'a, 'r> for MyJWTClaims {
    type Error = ();

    fn from_request(request: &'a Request<'r>) -> request::Outcome<Self, Self::Error> {
        let tokenStrings: Vec<_> = request.headers().get("Authorization").collect();

        match tokenStrings.len() {
            0 => Outcome::Failure((Status::Unauthorized, ())),
            1 => { 
                let client = reqwest::blocking::Client::new();
                let req = client.get("http://localhost:8081/api/users/secured/ping")
                                                .header("Authorization", tokenStrings[0])
                                                .send();
                match req {
                    Ok(res) => {
                        match res.error_for_status() {
                            Ok(_) => {
                                let verified: Result<Token<Header, MyJWTClaims, Unverified<'_>>, jwt::Error> = Token::parse_unverified(tokenStrings[0]);
                                match verified {
                                    Ok(token) => {
                                        let my_claims = MyJWTClaims {
                                            email: token.claims().email.clone(),
                                            role: token.claims().role.clone(), 
                                            exp: token.claims().exp,
                                        };
                                        println!("{}", my_claims.role);
                                        Outcome::Success(my_claims)
                                    }
                                    Err(_) => Outcome::Failure((Status::Unauthorized, ()))
                                }
                            },
                            Err(_) => Outcome::Failure((Status::Unauthorized, ()))
                        }
                    },
                    Err(_) => Outcome::Failure((Status::Unauthorized, ()))
                }
            },
            _ => Outcome::Failure((Status::Unauthorized, ())),
        }
    }
}