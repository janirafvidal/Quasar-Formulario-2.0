<template>
  <q-page>
    <div class="row justify-center">
      <q-card class="full-width">
        <q-card-section class="text-center">
          <q-avatar size="150px">
            <img
              src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/User_with_smile.svg/1024px-User_with_smile.svg.png"
            />
          </q-avatar>
          <q-input
            rounded
            v-model="form.nombre"
            label="Nombre"
            class="q-my-sm"
          />
          <q-input
            rounded
            v-model="form.apellidos"
            label="Apellidos"
            class="q-my-sm"
          />
          <q-input
            rounded
            v-model="form.email"
            label="Correo electrónico"
            hint="tucorreo@mail.com "
            class="q-my-sm"
          />
          <q-input
            rounded
            v-model="form.telefono"
            label="Teléfono"
            mask="(+##) ### ## ## ##"
            hint="Ejemplo: +34 600 00 00 00 "
            class="q-my-sm"
          />

          <q-input
            rounded
            v-model="form.fechaNacimiento"
            label="Fecha de nacimiento"
            mask="##/##/####"
            class="q-my-sm"
          />
          <q-select
            rounded
            v-model="form.genero"
            :options="opciones"
            label="Género"
            class="q-my-sm"
          />
          <q-btn
            color="primary"
            icon="send"
            label="Regístrame"
            class="q-my-sm"
            @click="send"
            id="btn"
          />
        </q-card-section>
      </q-card>
    </div>
  </q-page>
</template>

<script>
export default {
  data() {
    return {
      opciones: ["Masculino", "Femenino", "Prefiero no decirlo"],
      form: {
        nombre: "",
        apellidos: "",
        email: "",
        telefono: "",
        fechaNacimiento: "",
        genero: "",
      },
    };
  },
  methods: {
    send() {
      window.psgo
        .call("app.formulario.send", this.form)
        .then((msg) => {
          this.$router.push({
            name: "greeting",
            params: { msg: msg },
          });
        })
        .catch((err) => {
          console.log(err);
        });
    },
  },
};
</script>
