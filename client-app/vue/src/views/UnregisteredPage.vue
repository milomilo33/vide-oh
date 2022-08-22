<template>
  <div>
    <b-navbar fixed="top" toggleable="lg" type="light" variant="light">
    <b-navbar-brand href="/">Home Page</b-navbar-brand>
    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
    <b-collapse id="nav-collapse" is-nav>    
      <b-navbar-nav>

      </b-navbar-nav>  

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item right href="/Login">Login</b-nav-item>
        <b-nav-item right href="/Register">Register</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
    </b-navbar>

    <router-view>

    </router-view>
  </div>
</template>
