const API_URL = 'http://local.detik.com:8000'
const LOGIN_URL = API_URL + '/login'

export default {
    login(context, creds, redirect) {
        context.$http.post(LOGIN_URL, creds).then(response => {
            localStorage.setItem('id_token', response.body.id_token)
            if (redirect) {
                context.$router.replace(redirect)
            }
        }, response => {
            context.error = response.statusText
        })
    },

    logout(context) {
        localStorage.removeItem('id_token')
        context.$router.replace('/home')
    },

    isAuthenticated() {
        var jwt = localStorage.getItem('id_token')
        if (jwt) {
            return true
        }
        return false
    },

    getAuthHeader() {
        return {
            'Authorization': 'Bearer ' + localStorage.getItem('id_token')
        }
    }

}