<template>
    <v-app>
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
//import axios from "axios";
//import router from "../router";

export default {
    name: "user-page",
    data() {
       
    },
    methods: {
       
    }
};
</script>

<style scoped>

</style>
