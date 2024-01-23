import { createRouter, createWebHistory } from "vue-router"

import useAuthStore from "@/store/auth"

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: "/",
			name: "index",
			component: () => import("../views/index.vue"),
		},
		{
			path: "/login",
			name: "login",
			component: () => import("../views/login.vue"),
		},
		{
			path: "/task",
			name: "task",
			component: () => import("../views/task.vue"),
		},
		{
			path: "/task/create",
			name: "task-create",
			component: () => import("../views/create-task.vue"),
		},
		{
			path: "/:pathMatch(.*)*",
			name: "not-found",
			component: import("../views/not-found.vue"),
		},
	],
})

router.beforeEach((to, _) => {
	const authStore = useAuthStore()
	const routeName = to.name as string
	const publicRoutes = ["index", "login", "not-found"]

	if (publicRoutes.includes(routeName)) {
		if (authStore.authenticated && routeName == "login") {
			return { name: "task" }
		}
		return true
	} else {
		if (authStore.authenticated) {
			return true
		} else {
			setTimeout(() => alert("rectricted route"), 0)
			return { name: "login" }
		}
	}
})

export default router
