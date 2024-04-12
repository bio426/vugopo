<script setup lang="ts">
import { reactive, ref } from "vue"
import { useRouter } from "vue-router"

import authService from "@/service/auth"
import useAuthStore from "@/store/auth"

const router = useRouter()
const authStore = useAuthStore()

const form = reactive({ username: "", password: "" })
const loading = ref(false)

async function login() {
    loading.value = true
    const user = await authService.login({
        username: form.username,
        password: form.password,
    })
    loading.value = false
    authStore.setUser(user)
    router.push({ name: "dashboard" })
}
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Login</h1>
            <form @submit.prevent="login">
                <label class="form-control w-full mb-4">
                    <div class="label">
                        <span class="label-text">Username</span>
                    </div>
                    <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                </label>
                <label class="form-control w-full mb-8">
                    <div class="label">
                        <span class="label-text">Password</span>
                    </div>
                    <input class="input input-bordered w-full" type="password" required v-model="form.password" />
                </label>
                <button class="btn btn-block btn-primary" type="submit" :disabled="loading">
                    <span class="loading loading-spinner" v-if="loading" />
                    Log In
                </button>
            </form>
        </div>
    </div>
</template>
