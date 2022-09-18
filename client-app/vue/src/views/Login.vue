<template>
<div class="justify-content-center login">
  <b-alert v-model="showSuccessAlert" dismissible fade variant="danger">
      Bad credentials.
    </b-alert>
  <b-card title="Login">
    <b-form>
      <b-form-input
        id="input-1"
        v-model="email"
        placeholder="E-mail"
        required
      >
      </b-form-input>

      <b-form-input
        id="input-2"
        v-model="password"
        placeholder="Password"
        required
        type="password"
      >
      </b-form-input>
      <div class="mt-2">
        <b-button variant="primary" type="button" v-on:click="login()"
          >Login</b-button
        >
      </div>
    </b-form>
  </b-card>
</div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      email: "",
      password: "",
      showSuccessAlert: false,
    };
  },

  methods: {
    login() {
      var _this = this;
      if (this.email.trim() == "" && this.password.trim() == "") {
        _this.showSuccessAlert = true;
        return;
      }
      this.axios
        .post("api/users/login", {
          email: this.email,
          password: this.password,
        }, { withCredentials: true })
        .then((response) => {
          sessionStorage.setItem("token", response.data.token);
          this.findUserRole();
        })
          .catch((error) => {
          console.log(error);
          _this.showSuccessAlert = true;
        });
    },

    findUserRole() {
      var userRole = JSON.parse(
        atob(sessionStorage.getItem("token").split(".")[1])
      ).role;
      if (userRole == "Administrator") {
        this.$router.push("AdministratorPage");
      }
      if (userRole == "RegisteredUser") {
        this.$router.push("RegisteredPage");
      }
      if (userRole == "SupportUser") {
        this.$router.push("SupportPage");
      }
    },
  },
}
</script>

<style scoped>
.login {
  max-width: 40rem;
  background-color: #ffffff;
  margin: auto;
  margin-top: 100px;
  margin-bottom: 200px;
  padding: 20px;
}
</style>