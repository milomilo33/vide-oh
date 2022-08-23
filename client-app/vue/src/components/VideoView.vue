<template>
    <div>
        <h1>{{ video.title }}</h1>
        <br>
        <video controls style="max-width: 50%;">
            <source :src="'/api/videos/video-stream/' + video.filename" type="video/mp4" />
        </video>
        <br>

        <p>{{ video.description }}</p>
        <p><strong>Uploader:</strong> {{ video.ownerEmail }}</p>

        <b-container v-if="role !== 'UnregisteredUser'">

        <p v-if="total_rating !== 0">Rated {{ total_rating }}/5</p>

        <b-row>
            <b-col sm="3"></b-col>
            <b-col>
            <b-form-input id="input_rating"
                name="input_rating"
                placeholder="Rating (1-5)"
                class="mb-2 mr-sm-2 mb-sm-0"
                type="number"
                min="1"
                max="5"
                v-model="input_rating">
            </b-form-input>
            </b-col>
            <b-col>
            <b-button @click="rateVideo" variant="primary">Rate</b-button>
            </b-col>
            <b-col>
            <small v-if="your_rating !== 0"> Your current rating is {{ your_rating }}</small>
            </b-col>
            <b-col sm="3"></b-col>
        </b-row>
        <br>
        <br>

        <b-button @click="reportVideo" variant="warning">Report video</b-button>
        <br>
        <br>

        <h2>Comments</h2>

        <b-form-textarea
            id="input_comment"
            v-model="input_comment"
            placeholder="Your comment..."
            rows="3"
            max-rows="6"
        ></b-form-textarea>
        <br>
        <b-button @click="postComment" variant="primary">Post new comment</b-button>
        <br>
        <br>

        <div v-for="comment in comments" :key="comment.id">
            <b-card :title="comment.owner_email" :sub-title="comment.posted_at">
                <b-card-text>
                    {{ comment.body }}
                </b-card-text>
                <b-button v-if="role === 'Administrator' || (role === 'RegisteredUser' && comment.owner_email === current_email)" @click="deleteComment(comment.id)" variant="danger">Delete</b-button>
                <b-button @click="reportComment(comment.id)" variant="warning">Report comment</b-button>
            </b-card>
        </div>

        </b-container>

        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>{{ success_message }}</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>

        <b-modal ref="error-modal" hide-footer title="Error">
            <div class="d-block text-center">
                <p>{{ error_message }}</p>
            </div>
            <b-button class="mt-3" variant="outline-danger" block @click="hideErrorModal">Close</b-button>
        </b-modal>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                video: {},
                role: "",
                current_email: "",
                input_comment: "",
                comments: [],
                total_rating: 0,
                your_rating: 0,
                input_rating: "",
                success_message: "",
                error_message: "",
            };
        },

        methods: {
            rateVideo() {
                let body = {
                    rating_owner_email: this.current_email,
                    rating: Number(this.input_rating),
                    rating_video_id: this.video.ID,
                };

                this.axios.post(`/api/comments/ratings`, body , {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.success_message = "Successfully rated video!";
                    this.getTotalRating();
                    this.getYourRating();
                    this.showSuccessModal();
                })
                .catch(error => {
                    this.error_message = "Could not rate video.";
                    this.showErrorModal();
                    console.log(error);
                });
            },

            reportVideo() {
                this.axios.get(`/api/videos/report-video/${this.video.ID}` , {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.success_message = "Successfully reported video!";
                    this.showSuccessModal();
                })
                .catch(error => {
                    this.error_message = "Could not report video.";
                    this.showErrorModal();
                    console.log(error);
                });
            },

            postComment() {
                let body = {
                    owner_email: this.current_email,
                    body: this.input_comment,
                    video_id: this.video.ID,
                };

                this.axios.post(`/api/comments/comments`, body , {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.success_message = "Successfully posted a new comment!";
                    this.getComments();
                    this.showSuccessModal();
                })
                .catch(error => {
                    this.error_message = "Could not post comment.";
                    this.showErrorModal();
                    console.log(error);
                });
            },

            deleteComment(id) {
                this.axios.get(`/api/comments/comments/delete/${id}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.success_message = "Successfully deleted comment!";
                    this.getComments();
                    this.showSuccessModal();
                })
                .catch(error => {
                    this.error_message = "Could not delete comment.";
                    this.showErrorModal();
                    console.log(error);
                });
            },

            reportComment(id) {
                this.axios.get(`/api/comments/comments/report/${id}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.success_message = "Successfully reported comment!";
                    this.showSuccessModal();
                })
                .catch(error => {
                    this.error_message = "Could not report comment.";
                    this.showErrorModal();
                    console.log(error);
                });
            },

            getTotalRating() {
                this.axios.get(`/api/comments/ratings/total/${this.video.ID}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.total_rating = response.data;
                })
                .catch(error => {
                    this.total_rating = 0;
                    console.log(error);
                });
            },

            getYourRating() {
                this.axios.get(`/api/comments/ratings/user/${this.current_email}/${this.video.ID}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.your_rating = response.data;
                })
                .catch(error => {
                    this.your_rating = 0;
                    console.log(error);
                });
            },

            getComments() {
                this.axios.get(`/api/comments/comments/${this.video.ID}`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.comments = response.data;
                    this.comments = this.comments.sort(function(x, y){
                        return new Date(y.posted_at) - new Date(x.posted_at);
                    })
                })
                .catch(error => {
                    this.comments = [];
                    console.log(error);
                });
            },
        
            showSuccessModal() {
                this.$refs['success-modal'].show()
            },

            hideSuccessModal() {
                this.$refs['success-modal'].hide()
            },

            showErrorModal() {
                this.$refs['error-modal'].show()
            },

            hideErrorModal() {
                this.$refs['error-modal'].hide()
            },
        },

        mounted() {
            let tokenString = sessionStorage.getItem('token');
            if (tokenString) {
                let token = JSON.parse(atob(tokenString.split('.')[1]));
                this.role = token.role;
                this.current_email = token.email;
            } else { 
                this.role = "UnregisteredUser";
                this.current_email = "";
            }

            this.video = this.$route.params.video;

            if (this.role !== "UnregisteredUser") {
                this.getComments();
                this.getTotalRating();
                this.getYourRating();
            }
        }
    }
</script>