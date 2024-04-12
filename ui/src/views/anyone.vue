<script setup lang="ts">
import { } from "vue"

import alertService from "@/service/alert"
import Header from "@/components/Header.vue";

let events: EventSource

function initEventSub() {
    events = new EventSource("http://localhost:1323/api/alert/sub", {
        withCredentials: true
    })
    events.addEventListener("open", () => {
        console.log("connected to events")
    })
    events.addEventListener("message", e => {
        console.log(e)
    })
}
initEventSub()

async function dispatchEvent() {
    await alertService.pub()
}
</script>

<template>
    <div class="relative w-screen min-h-screen ">
        <div class="w-11/12 mx-auto">
            <Header title="Anyone" />
            <p>
                Lorem ipsum, dolor sit amet consectetur adipisicing elit. Consequatur sed adipisci nihil vero amet aut,
                perferendis voluptate? Dignissimos corrupti ad qui commodi mollitia perspiciatis non, repellat fugiat error
                sunt vel!
            </p>
            <button class="btn" @click="dispatchEvent">Publish</button>
        </div>
    </div>
</template>
