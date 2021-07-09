<template>
  <q-page>
    <div class="row justify-center">
      <q-card class="col-5">
        <q-card-section class="text-center">
          <q-avatar size="150px">
            <img
              src="https://upload.wikimedia.org/wikipedia/commons/thumb/5/5f/User_with_smile.svg/1024px-User_with_smile.svg.png"
            />
          </q-avatar>
          <div>{{ greeting }}</div>
        </q-card-section>
      </q-card>
    </div>
  </q-page>
</template>

<script>
export default {
  data() {
    return {
      greeting: "",
    };
  },
  created() {
    this.greeting = this.$route.params.msg;
  },
};
</script>
