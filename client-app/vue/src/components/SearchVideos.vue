<template>
    <div>
        <h1>Videos</h1>
        <br>
        <div style="width:100%;text-align:center;">
            <b-form inline style="display: inline-block;">
                <!-- <label>Keyword:</label> -->
                <b-form-input id="keyword"
                            name="keyword"
                            placeholder="Keyword..."
                            class="mb-2 mr-sm-2 mb-sm-0"
                            v-model="keyword">
                </b-form-input>
                <b-button @click="searchVideos" class="mb-2 mr-sm-2 mb-sm-0">Search</b-button>
            </b-form>
        </div>
        <br>

        <b-container class="bv-example-row">
            <b-row>
                <div v-for="video in videos" :key="video.ID">
                    <b-card
                        :title="video.title"
                        :img-src="'/api/videos/static/' + video.filename + '.jpg'"
                        img-alt="Thumbnail"
                        img-top
                        tag="article"
                        style="max-width: 20rem;"
                        class="mb-2"
                    >
                        <b-card-text>
                            {{ video.description }}
                        </b-card-text>

                        <b-card-text>
                            Posted by: {{ video.ownerEmail }}
                        </b-card-text>

                        <b-button @click="openVideoView(video)" variant="primary">Open</b-button>
                        <b-button v-if="role === 'Administrator' || (role === 'RegisteredUser' && video.ownerEmail === current_email)" @click="deleteVideo(video.ID)" variant="danger">Delete</b-button>
                    </b-card>
                </div>
            </b-row>
        </b-container>

        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>Video successfully deleted.</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                keyword: "",
                videos: [],
                role: "",
                current_email: "",
            };
        },

        methods: {
            openVideoView(video) {
                this.$router.push({
                    name: 'VideoView' + this.role,
                    params: { video }
                });
            },

            searchVideos() {
                this.axios.get(`/api/videos/search-videos?query=${this.keyword}`)
                .then((response) => {
                    this.videos = response.data;
                    console.log(response.data);
                })
                .catch(error => {
                    console.log(error);
                });
            },

            deleteVideo(id) {
                this.axios.get(`/api/videos/secured/delete-video/${id}`, {
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
            let tokenString = sessionStorage.getItem('token');
            if (tokenString) {
                let token = JSON.parse(atob(tokenString.split('.')[1]));
                this.role = token.role;
                this.current_email = token.email;
            } else { 
                this.role = "UnregisteredUser";
                this.current_email = "";
            }

            this.keyword = "";

            this.searchVideos();
        }
    }
</script>