import Vuelidate from 'vuelidate'

export default ({ app }) => {
  console.log("antes")
  app.use(Vuelidate)
  console.log("despues")
}