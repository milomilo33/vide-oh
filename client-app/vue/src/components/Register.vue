<template>
    <b-container>
        <h1>Registration</h1>
        <b-form>
            <b-form-input id="email"
                        name="email"
                        placeholder="Email"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="email">
            </b-form-input>
            <br>
            <b-form-input id="password"
                        name="password"
                        placeholder="Password"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        type="password"
                        v-model="password">
            </b-form-input>
            <br>
            <b-form-input id="name"
                        name="name"
                        placeholder="Name"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="name">
            </b-form-input>
            <br>
            <b-button @click="onSubmit" class="mb-2 mr-sm-2 mb-sm-0">Register</b-button>
        </b-form>

        <b-modal ref="error-modal" hide-footer title="Error">
            <div class="d-block text-center">
                <p>{{ this.errorMessage }}</p>
            </div>
            <b-button class="mt-3" variant="outline-danger" block @click="hideErrorModal">Close</b-button>
        </b-modal>
        
        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>Successfully registered.</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>

    </b-container>
</template>

<script>

    export default {
        data() {
            return {
                password: '',
                name: '',
                email: '',
                errorMessage: ''
            }
        },
        methods: {
            onSubmit() {
                let body = {
                    password: this.password,
                    name: this.name,
                    email: this.email
                };

                this.errorMessage = ""
                let reemail = /^\w+@[a-zA-Z_]+?\.[a-zA-Z]{2,3}$/;   
                if (body.name === "") {
                    this.errorMessage = "Name can not be empty";
                    this.showErrorModal();
                }
                if (body.password === "") {
                    this.errorMessage = "Passowrd can not be empty";
                    this.showErrorModal();
                }
                if (!reemail.test(body.email)) {
                    this.errorMessage = "Provide valid email";
                    this.showErrorModal();
                }

                if (this.errorMessage !== ""){
                    return
                }

                this.axios.post(`/api/users/user/register`, body)
                .then(() => {
                    this.showSuccessModal();
                    this.$router.push("/Login");
                })
                .catch(error => {
                    this.errorMessage = error.response.data;
                    this.showErrorModal();
                });
            },

            hideErrorModal() {
                this.$refs['error-modal'].hide()
            },

            hideSuccessModal() {
                this.$refs['success-modal'].hide()
            },

            showErrorModal() {
                this.$refs['error-modal'].show()
            },

            showSuccessModal() {
                this.$refs['success-modal'].show()
            },
        }
    }
</script>
