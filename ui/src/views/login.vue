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
	router.push({ name: "task" })
}
</script>

<template>
	<div>
		<form @submit.prevent="login">
			<label>
				Username
				<input type="text" required v-model="form.username" />
			</label>
			<label>
				Password
				<input type="password" required v-model="form.password" />
			</label>
			<button :aria-busy="loading" type="submit">Log In</button>
		</form>
	</div>
</template>
