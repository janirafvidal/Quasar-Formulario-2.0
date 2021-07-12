<template>
  <q-page>
    <h4 class="row justify-center">LISTADO DE USUARIOS REGISTRADOS</h4>
    <div class="row justify-center q-pa-md">
      <q-btn label="+ AÑADIR USUARIO" color="primary" @click="btn = true">
        <!-- <q-tooltip class="bg-accent">Añadir usuario</q-tooltip> -->
      </q-btn>

      <q-dialog v-model="btn">
        <q-card class="full-width">
          <q-card-section class="q-pt-none">
            <Form />
          </q-card-section>
        </q-card>
      </q-dialog>
    </div>

    <div class="justify-center row>
      <q-list bordered class="rounded-borders  col-8"">
        <q-expansion-item v-for="(user, id) in userList" :key="id">
          <template v-slot:header>
            <q-item-section avatar>
              <q-avatar>
                <img src="https://cdn.quasar.dev/img/boy-avatar.png" />
              </q-avatar>
            </q-item-section>

            <q-item-section> {{ user.name }} </q-item-section>
          </template>

          <q-card>
            <q-card-section>
              <p>{{ user.mail }}</p>
              <p>Edad: {{ user.age }}</p>
              <p>Teléfono de contacto: {{ user.phone }}</p>
              <p>Población: {{ user.address }}</p>
              <q-checkbox
                disable
                left-label
                v-model="user.isAdmin"
                label="Es administrador"
              />
            </q-card-section>
          </q-card>
        </q-expansion-item>

        <q-separator />
      </q-list>
    </div>
  </q-page>
</template>

<script>
import Form from "../components/Form.vue";

export default {
  data() {
    return {
      btn: false,
      userList: [],
    };
  },
  components: {
    Form,
  },
  methods: {
    getUser() {
      window.psgo
        .call("app.formulario.get_users")
        .then((msg) => {
          this.userList = msg;
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
  created() {
    this.getUser();
  },
};
</script>
