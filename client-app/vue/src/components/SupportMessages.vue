<template>
    <div>
        <h1>Tech Support</h1>
        <span>
            <label>Messages with: </label>
            <b-form-select
                id="selectUser"
                name="selectUser"
                placeholder="Select user"
                class="mb-2 mr-sm-2 mb-sm-0"
                v-model="selected_user_email"
                :options="user_email_options"
                @change="optionChanged">
            </b-form-select>
        </span>
        <br>

        <router-view :key="$route.fullPath">
        </router-view>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                selected_user_email: "",
                user_email_options: []
            };
        },

        methods: {
            getMessages() {
                this.axios.get(`/api/messages/secured/user-emails`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.user_email_options = response.data;
                })
                .catch(error => {
                    this.user_email_options = [];
                    console.log(error);
                });
            },

            optionChanged() {
                this.$router.push({
                    name: 'SupportUserMessages',
                    query: { owner_email: this.selected_user_email },
                });
            }
        },

        mounted() {
            this.getMessages();
        }
    }
</script>