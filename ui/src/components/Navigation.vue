<script setup lang="ts">
import { ref } from "vue"
import { useRouter } from "vue-router"

import authService from "@/service/auth"
import useAuthStore from "@/store/auth"

const router = useRouter()
const authStore = useAuthStore()
const detailEl = ref<HTMLDetailsElement>()

function closeDetail() {
	if (detailEl.value == undefined) return
	detailEl.value.open = false
}

async function logout() {
	await authService.logout()
	authStore.user = undefined
	router.push({ name: "login" })
}
</script>

<template>
	<nav>
		<ul>
			<li><strong>Vugopo</strong></li>
			<li>
				<details role="list" ref="detailEl">
					<summary aria-haspopup="listbox">≡</summary>
					<ul role="listbox">
						<li @click="closeDetail">
							<router-link :to="{ name: 'index' }">
								Index
							</router-link>
						</li>
						<template v-if="authStore.authenticated">
							<li>
								<hr />
							</li>
							<li @click="closeDetail">
								<router-link :to="{ name: 'task' }">
									My Tasks
								</router-link>
							</li>
						</template>
					</ul>
				</details>
			</li>
		</ul>
		<ul>
			<li v-if="authStore.user">
				<b>{{ authStore.user.username }}</b>
				<br />
				<a class="secondary" @click="logout">
					<small>
						<u>Logout</u>
					</small>
				</a>
			</li>
			<li v-else-if="!authStore.authenticated && $route.name != 'login'">
				<router-link :to="{ name: 'login' }">Login</router-link>
			</li>
		</ul>
	</nav>
</template>
