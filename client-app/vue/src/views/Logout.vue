<template>
  <p>redirecting...</p>
</template>

<script>

export default {
  name: "Logout",
  mounted() {
    sessionStorage.removeItem("token");
    this.$router.push("/");
  }
}
</script>
