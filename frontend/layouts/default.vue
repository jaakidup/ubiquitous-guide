<template>
  <div class="main container">
    <b-navbar class="navbar" variant="faded" type="light">
      <!-- <b-navbar-brand  v-b-tooltip.hover title="Go to the HOME Page"> -->

        <b-link class="link" to="/">
          EYE <img src="~/assets/eye.jpg" class="header-image d-inline-block align-top" alt="EYE" />
          DEA
        </b-link>
        <div class="spacer"></div>
        <div class="spacer"></div>
      <!-- </b-navbar-brand> -->
    </b-navbar>

    <div class="container">
      <nuxt />
    </div>
  </div>
</template>

<style>
html {
  font-family: "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI",
    Roboto, "Helvetica Neue", Arial, sans-serif;
  font-size: 16px;
  word-spacing: 1px;
  -ms-text-size-adjust: 100%;
  -webkit-text-size-adjust: 100%;
  -moz-osx-font-smoothing: grayscale;
  -webkit-font-smoothing: antialiased;
  box-sizing: border-box;


}
body {
  background: black;
  padding: 1em;
}

*,
*:before,
*:after {
  box-sizing: border-box;
  margin: 0;  
}


.container {
  max-width: 800px;
  /* margin: 0 auto; */
  min-height: 100vh;
  /* display: flex; */
  /* justify-content: center; */
  /* align-items: center; */
  /* text-align: center; */
}

.main {
  background: white;
  padding-top: 1px;
}


.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont,
    "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.navbar {
  margin-top: 2em;
  background: black;

}

.links {
  padding-top: 15px;
}
.link {
  text-decoration: none;
  color: white;
  padding-left: 30rem;
  padding-right: 10rem;
}
.link:hover {
  text-decoration: none;
  border: none;
  color: greenyellow;
  
}
.header-image {
  max-width: 30px;
}
b-navbar-brand {
  margin-left: auto;
  margin-right: auto;
  }
  .spacer {
    background: silver;
    width: 100%;
    margin-right: auto;
  }
</style>
