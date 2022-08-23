<template>
    <b-container>
        <h1>Profile</h1>
        <b-form>
            <b-form-input id="email"
                        name="email"
                        placeholder="Email"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="email"
                        disabled>
            </b-form-input>
            <br>
            <b-form-input id="name"
                        name="name"
                        placeholder="Name"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="name">
            </b-form-input>
            <br>
            <b-button @click="onSubmit" class="mb-2 mr-sm-2 mb-sm-0">Update</b-button>
        </b-form>

        <b-modal ref="error-modal" hide-footer title="Error">
            <div class="d-block text-center">
                <p>{{ this.errorMessage }}</p>
            </div>
            <b-button class="mt-3" variant="outline-danger" block @click="hideErrorModal">Close</b-button>
        </b-modal>
        
        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>Successfully updated profile.</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>

    </b-container>
</template>

<script>

    export default {
        data() {
            return {
                name: '',
                email: '',
                errorMessage: ''
            }
        },
        methods: {
            onSubmit() {
                this.errorMessage = ""
                if (this.name === "") {
                    this.errorMessage = "Name can not be empty";
                    this.showErrorModal();
                }

                if (this.errorMessage !== ""){
                    return
                }

                this.axios.get(`/api/users/secured/user/change-name?name=${this.name}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then(() => {
                    this.showSuccessModal();
                    this.getCurrentUser();
                })
                .catch(error => {
                    this.errorMessage = "Could not update profile.";
                    this.showErrorModal();
                });
            },

            getCurrentUser() {
                this.axios.get(`/api/users/secured/user/current`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.name = response.data.name;
                    this.email = response.data.email;
                })
                .catch(error => {
                    console.log(error);
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
        },
        mounted() {
            this.getCurrentUser();
        }
    }
</script>
