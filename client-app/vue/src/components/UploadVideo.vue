<template>
    <b-container>
        <h1>Video upload</h1>
        <b-form>
            <b-form-input id="title"
                        name="title"
                        placeholder="Title"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="title">
            </b-form-input>
            <br>
            <b-form-input id="description"
                        name="description"
                        placeholder="Description"
                        class="mb-2 mr-sm-2 mb-sm-0"
                        v-model="description">
            </b-form-input>
            <br>
            <b-form-file
                v-model="video_file"
                :state="Boolean(video_file)"
                placeholder="Choose an MP4 video file or drop it here..."
                drop-placeholder="Drop file here..."
                accept=".mp4"
                ></b-form-file>
            <br>
            <br>
            <b-button @click="onSubmit" class="mb-2 mr-sm-2 mb-sm-0">Upload</b-button>
        </b-form>

        <b-modal ref="error-modal" hide-footer title="Error">
            <div class="d-block text-center">
                <p>{{ this.errorMessage }}</p>
            </div>
            <b-button class="mt-3" variant="outline-danger" block @click="hideErrorModal">Close</b-button>
        </b-modal>
        
        <b-modal ref="success-modal" hide-footer title="Success">
            <div class="d-block text-center">
                <p>Successfully uploaded video!</p>
            </div>
            <b-button class="mt-3" variant="outline-success" block @click="hideSuccessModal">Close</b-button>
        </b-modal>

    </b-container>
</template>

<script>

    export default {
        data() {
            return {
                title: '',
                description: '',
                video_file: null,
                errorMessage: ''
            }
        },
        methods: {
            onSubmit() {
                this.errorMessage = ""
                if (this.title === "") {
                    this.errorMessage = "Title can not be empty";
                    this.showErrorModal();
                }
                if (this.description === "") {
                    this.errorMessage = "Description can not be empty";
                    this.showErrorModal();
                }
                if (!this.video_file) {
                    this.errorMessage = "File can not be empty";
                    this.showErrorModal();
                }

                if (this.errorMessage !== ""){
                    return
                }

                let formData = new FormData();
                formData.append("file", this.video_file);

                this.axios.post(`http://localhost:8082/api/videos/secured/upload-video?title=${this.title}&description=${this.description}`, formData, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                            'Content-Type': 'multipart/form-data',
                        },         maxContentLength: 100000000,
        maxBodyLength: 1000000000
                    })
                .then(() => {
                    this.showSuccessModal();
                    this.$router.push("/RegisteredPage");
                })
                .catch(error => {
                    console.log(error);
                    this.errorMessage = "Failed to upload video."
                    this.showErrorModal();
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
        }
    }
</script>
