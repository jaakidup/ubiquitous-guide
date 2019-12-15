<template>
  <div>
    User
    <br>
    {{ params }}
  </div>
</template>

<script>
export default {
    asyncData({params}) {
      return {
        params: params
      }
    }
}
</script>

<style>

</style>