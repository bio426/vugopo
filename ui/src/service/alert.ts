import base from "./base"

const path = "alert"

export default {
    async pub() {
        const res = await base.post(path + "/pub")

        return res.status
    },
}
