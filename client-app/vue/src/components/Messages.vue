<template>
    <div>
        <br>
        <br>
        <h2 v-if="role === 'SupportUser'">{{ owner_email }}</h2>
        <br>

        <b-container>

        <div v-for="message in messages" :key="message.ID">
            <b-card :title="message.sentByUser ? message.ownerEmail : 'Tech Support'" :sub-title="message.date">
                <b-card-text>
                    {{ message.content }}
                </b-card-text>
            </b-card>
        </div>
        <br>

        <b-form-textarea
            id="input_message"
            v-model="input_message"
            placeholder="Your message..."
            rows="3"
            max-rows="6"
        ></b-form-textarea>
        <br>
        <b-button @click="sendMessage" variant="primary">Send message</b-button>
        <br>
        <br>
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
                owner_email: "",
                role: "",
                token_email: "",
                input_message: "",
                messages: [],
                success_message: "",
                error_message: "",
                socket: {},
                selected_user_email: "",

            };
        },

        methods: {
            sendMessage() {
                let body = {
                    message: this.input_message,
                    token: sessionStorage.getItem('token')
                };

                this.socket.send(JSON.stringify(body));
            },

            getMessages() {
                this.axios.get(`/api/messages/secured/${this.owner_email}/all`, {
                        headers: {
                            Authorization: sessionStorage.getItem('token'),
                        },
                    })
                .then((response) => {
                    this.messages = response.data;
                    this.messages = this.messages.sort(function(x, y){
                        return new Date(x.date) - new Date(y.date);
                    })
                })
                .catch(error => {
                    this.messages = [];
                    console.log(error);
                });
            },

            setUpSockets() {
                // this.socket = new SockJS(`ws://localhost:8084/api/messages/${this.owner_email}/ws`);
                // this.stompClient = Stomp.over(this.socket);
                // this.stompClient.connect(
                //     {Authorization: "Bearer " + sessionStorage.getItem("token")},
                //     frame => {
                //     // console.log(frame);
                //     this.stompClient.subscribe("/topic/alarms", tick => {
                //         this.makeToast(false, tick.body);
                //         // this.received_messages.push(JSON.parse(tick.body).content);
                //     });
                //     },
                //     error => {
                //     console.log(error);
                //     this.connected = false;
                //     }
                // );

                this.socket = new WebSocket(`ws://localhost:8084/api/messages/${this.owner_email}/ws?token=${sessionStorage.getItem('token')}`);
                let _this = this;
                this.socket.onmessage = function (msg) {
                    _this.messages.push(JSON.parse(msg.data));
                    // console.log(msg);
                };
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

            this.owner_email = this.$route.query.owner_email;

            if (this.role !== "UnregisteredUser") {
                this.getMessages();
                this.setUpSockets();
            }
        },

        destroyed() {
            this.socket.close();
            console.log("socket closed");
        }
    }
</script>