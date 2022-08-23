<template>
    <div>
        <h1>Reported videos</h1>
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
                        <b-button @click="deleteVideoAndBlockUser(video)" variant="danger">Delete video & block user</b-button>
                    </b-card>
                </div>
            </b-row>
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
                videos: [],
                role: "",
                current_email: "",
            };
        },

        methods: {
            openVideoView(video) {
                this.$router.push({
                    name: 'VideoViewAdministrator',
                    params: { video }
                });
            },

            getReportedVideos() {
                this.axios.get(`/api/videos/secured/all-reported-videos`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.videos = response.data;
                    console.log(response.data);
                })
                .catch(error => {
                    console.log(error);
                });
            },

            deleteVideoAndBlockUser(video) {
                this.axios.get(`/api/videos/secured/delete-video/${video.ID}`, {
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

                this.axios.get(`/api/users/secured/block/${video.ownerEmail}`, {
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
            this.getReportedVideos();
        }
    }
</script>