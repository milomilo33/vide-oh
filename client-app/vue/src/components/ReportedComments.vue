<template>
    <div>
        <h1>Reported comments</h1>
        <br>

        <b-container class="bv-example-row">
            <div v-for="comment in comments" :key="comment.id">
                <b-card :title="comment.owner_email" :sub-title="comment.posted_at">
                    <b-card-text>
                        {{ comment.body }}
                    </b-card-text>
                    <b-button @click="deleteCommentAndBlockUser(comment)" variant="danger">Delete comment & block user</b-button>
                </b-card>
            </div>
        </b-container>

        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>Success.</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                comments: [],
            };
        },

        methods: {
            getReportedComments() {
                this.axios.get(`/api/comments/comments/reported`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.comments = response.data;
                    console.log(response.data);
                })
                .catch(error => {
                    console.log(error);
                });
            },

            deleteCommentAndBlockUser(comment) {
                this.axios.get(`/api/comments/comments/delete/${comment.id}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.showSuccessModal();
                })
                .catch(error => {
                    console.log(error);
                });

                this.axios.get(`/api/users/secured/block/${comment.owner_email}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.showSuccessModal();
                })
                .catch(error => {
                    console.log(error);
                });
            },
        
            showSuccessModal() {
                this.$refs['success-modal'].show()
            },

            hideSuccessModal() {
                this.$refs['success-modal'].hide()
            },
        },

        mounted() {
            this.getReportedComments();
        }
    }
</script>