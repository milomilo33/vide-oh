<template>
  <div>
    <b-navbar fixed="top" toggleable="lg" type="light" variant="light">
    <b-navbar-brand href="/SupportPage/SupportMessages">Home Page</b-navbar-brand>
    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
    <b-collapse id="nav-collapse" is-nav>    
      <b-navbar-nav>
      </b-navbar-nav>  

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item-dropdown right>
          <!-- Using 'button-content' slot -->
          <template #button-content>
            <em>{{ current_name }} ({{ current_email }})</em>
          </template>
          <b-dropdown-item href="/SupportPage/Profile">Profile</b-dropdown-item>
          <b-dropdown-item href="/Logout">Log out</b-dropdown-item>
        </b-nav-item-dropdown>
      </b-navbar-nav>
    </b-collapse>
    </b-navbar>

    <router-view>

    </router-view>
  </div>
</template>

<script>

export default {
  data() {
      return {
          current_name: "",
          current_email: "",
      };
  },

  mounted() {
    this.axios.get(`/api/users/secured/user/current`, {
            headers: {
                Authorization: sessionStorage.getItem('token'),
            },
        })
    .then((response) => {
        this.current_name = response.data.name;
        this.current_email = response.data.email;
    })
    .catch(error => {
        this.current_name = "ERROR";
        this.current_email = "ERROR";
        console.log(error);
    });
  }
}
</script>
