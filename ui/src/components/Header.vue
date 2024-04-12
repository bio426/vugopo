<script setup lang="ts">
import {} from "vue"
import { useRouter } from "vue-router"
import { Bars3Icon, HomeIcon, UsersIcon } from "@heroicons/vue/24/solid"

import authService from "@/service/auth"
import useAuthStore from "@/store/auth"

const props = defineProps<{ title: string }>()

const router = useRouter()
const authStore = useAuthStore()

const u4: boolean = 3

async function logout() {
	await authService.logout()
	authStore.clearUser()
	router.push({ name: "login" })
}
</script>

<template>
	<div class="navbar bg-base-100">
		<div class="navbar-start">
			<div class="dropdown">
				<label tabindex="0" class="btn btn-ghost btn-circle">
					<Bars3Icon class="w-6" />
				</label>
				<ul
					tabindex="0"
					class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					<li>
						<router-link
							active-class="active"
							:to="{ name: 'dashboard' }"
						>
							<HomeIcon class="w-4" />
							Dashboard
						</router-link>
					</li>
					<li v-if="authStore.user?.role == 'super'">
						<router-link
							active-class="active"
							:to="{ name: 'user' }"
						>
							<UsersIcon class="w-4" />
							Users
						</router-link>
					</li>
					<!-- <li> -->
					<!-- 	<router-link -->
					<!-- 		active-class="active" -->
					<!-- 		:to="{ name: 'pos' }" -->
					<!-- 	> -->
					<!-- 		<BanknotesIcon class="w-4" /> -->
					<!-- 		Pos -->
					<!-- 	</router-link> -->
					<!-- </li> -->
					<!-- <li> -->
					<!-- 	<details open> -->
					<!-- 		<summary> -->
					<!-- 			<CubeIcon class="w-4" /> -->
					<!-- 			Products -->
					<!-- 		</summary> -->
					<!-- 		<ul> -->
					<!-- 			<li> -->
					<!-- 				<router-link -->
					<!-- 					active-class="active" -->
					<!-- 					:to="{ name: 'products-list' }" -->
					<!-- 				> -->
					<!-- 					List -->
					<!-- 				</router-link> -->
					<!-- 			</li> -->
					<!-- 			<li> -->
					<!-- 				<router-link -->
					<!-- 					active-class="active" -->
					<!-- 					:to="{ name: 'products-create' }" -->
					<!-- 				> -->
					<!-- 					Create -->
					<!-- 				</router-link> -->
					<!-- 			</li> -->
					<!-- 		</ul> -->
					<!-- 	</details> -->
					<!-- </li> -->
				</ul>
			</div>
		</div>
		<div class="navbar-center">
			<h1 class="py-4 text-2xl font-bold text-center">
				{{ props.title }}
			</h1>
		</div>
		<div class="navbar-end">
			<span>{{ authStore.user?.role }}</span>
			<div class="dropdown dropdown-end">
				<div
					tabindex="0"
					role="button"
					class="btn btn-ghost btn-circle avatar"
				>
					<div class="w-10 rounded-full">
						<img
							src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"
						/>
					</div>
				</div>
				<ul
					tabindex="0"
					class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					<li>
						<button @click="logout">Logout</button>
					</li>
				</ul>
			</div>
		</div>
	</div>
</template>
