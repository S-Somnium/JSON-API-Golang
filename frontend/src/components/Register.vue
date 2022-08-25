<template>
    <v-app id="inspire">
        <v-content>
            <v-container fluid fill-height>
                <v-layout align-center justify-center>
                    <v-flex xs12 sm8 md4>
                        <v-card class="elevation-12 card">

                            <v-form ref="form">
                                <v-text-field :counter="20" label="Name" v-model="name">
                                </v-text-field>

                                <v-text-field :counter="254" label="E-mail" v-model="email">
                                </v-text-field>

                                <v-text-field :counter="254" label="Password" v-model="password" :type="'password'">
                                </v-text-field>

                                <v-slider v-model="age" color="orange" label="Age" hint="Be honest" min="19" max="100"
                                    thumb-label></v-slider>


                                <v-btn class="mr-4 register" @click="register()">
                                    Create
                                </v-btn>

                            </v-form>
                        </v-card>

                    </v-flex>
                </v-layout>
            </v-container>
        </v-content>
    </v-app>
</template>


<script>
import axios from "axios";
import router from "../router";

export default {
    name: "register-page",
    data() {
        return {
            name: '',
            email: '',
            password: '',
            age: ''
        }
    },
    methods: {
        register() {
            const { name, email, password, age } = this
            axios
                .post('http://localhost:8080/register', { name, email, password, age: parseInt(age) })
                .then(response => {
                    alert(response.data.message)
                    router.push({ name: 'login' })
                })
                .catch(error => alert(error.response.data.error))
        }
    }
};
</script>

<style scoped>
.card {
    padding: 10px;
}
</style>
