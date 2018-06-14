<template>
  <div id="app">
    <transition name="fade"
		            mode="out-in">
			<router-view></router-view>
		</transition>
  </div>
</template>

<script>
export default {
  name: 'app',
	components: {
	}
}
</script>

<style>
#app {
  font-family: "Lato","proxima-nova","Helvetica Neue",Arial,sans-serif;
}
html, body {
  margin: 0px;
}
</style>
