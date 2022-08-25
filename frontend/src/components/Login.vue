<template>
    <v-app>
        <v-content>
            <v-container fluid fill-height>
                <v-layout align-center justify-center>
                    <v-flex xs12 sm8 md4>
                        <v-card class="elevation-12">
                            <v-toolbar dark color="primary">
                                <v-toolbar-title>Login form</v-toolbar-title>
                            </v-toolbar>
                            <v-card-text>
                                <v-form>
                                    <v-text-field v-model="email" label="Login" type="text"></v-text-field>
                                    <v-text-field v-model="password" label="Password" type="password">
                                    </v-text-field>
                                </v-form>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn color="white" class="mr-4">
                                    <router-link to="/register">register</router-link>
                                </v-btn>
                                <v-spacer></v-spacer>
                                <v-btn color="primary" @click="login()">Login</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-content>
    </v-app>
</template>

<script>
import axios from "axios";
//import router from "../router";
const test = () => {
    let token = localStorage.getItem('token')
    axios.get('http://localhost:8080/user/', {
        headers: {
            Authorization: `Bearer ${token}`, "Content-Type": "application/json"
        }
    })
        .then(response => {
            console.log(response)
        })
        .catch(error => console.log(error))
}
export default {
    name: "login-page",
    data() {
        return {
            email: '',
            password: '',
        }
    },
    methods: {
        login() {
            const { email, password } = this
            console.log(email, password)
            axios
                .post('http://localhost:8080/login', { email, password })
                .then(response => {
                    console.log(response)
                    localStorage.setItem("token", response.data.token)
                    test()
                    //router.push({name:'login'})
                })
                .catch(error => alert(error.response.data.error))
        }
    }
};
</script>

<style>
</style>
