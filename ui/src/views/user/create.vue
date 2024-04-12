<script setup lang="ts">
import { ref, reactive } from "vue"
import { useRouter } from "vue-router";

import authService from "@/service/auth"

const router = useRouter()

const loading = ref(false)

const form = reactive({
    username: "",
    password: "",
    role: "",
})

const roleOpts = ["owner", "employee"]

async function create() {
    loading.value = true
    await authService.register({ username: form.username, password: form.password, role: form.role })
    loading.value = false
    router.push({ name: "user" })
}
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <h1 class="py-8 text-2xl font-bold text-center">Create User</h1>
            <form @submit.prevent="create" autocomplete="off">
                <label class="form-control w-full mb-4">
                    <div class="label">
                        <span class="label-text">Username</span>
                    </div>
                    <input class="input input-bordered w-full" type="text" required v-model="form.username" />
                </label>
                <label class="form-control w-full mb-4">
                    <div class="label">
                        <span class="label-text">Password</span>
                    </div>
                    <input class="input input-bordered w-full" type="text" required v-model="form.password" />
                </label>
                <label class="form-control w-full mb-8">
                    <div class="label">
                        <span class="label-text">Role</span>
                    </div>
                    <select class="select select-bordered w-full" required v-model="form.role">
                        <option disabled selected>---</option>
                        <option :value="role" v-for="role in roleOpts">{{ role }}</option>
                    </select>
                </label>

                <button class="btn btn-block btn-primary" type="submit" :disabled="loading">
                    <span class="loading loading-spinner" v-if="loading" />
                    Create
                </button>
            </form>
        </div>
    </div>
</template>
