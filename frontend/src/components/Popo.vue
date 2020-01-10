<template>
  <section id="app">
    <section vcloak id="window">
      <label id="gt" for="query" class="unselectable">></label>
      <input 
        id="query"
        placeholder=". . ."
        autofocus 
        autocomplete="off" 
        v-model="query" 
        v-on:input="search"
        @keydown="navigate"
        >

        <ul id="searchResults" class="unselectable">
          <li 
            v-for="(searchResult, index) in searchResults" 
            class="searchResult"
            :class="{ selectedResult: navIdx === index }"
            v-on:click="selectResult(index)"
            @mouseover="highlightResult(index)"
            >
            <img class="file-icon" src="./assets/images/icons8-file-50-4.png"></img>
            <div class="file">{{searchResult.Name}}
              <footer class="filepath">{{searchResult.Path}}</footer>
            </div>
          </li>
        </ul>

    </section>
  </section>
</template>

<script>
const UP_ARROW_KEY = 38;
const DOWN_ARROW_KEY = 40;
const ENTER_KEY = 13;

export default {
  // app initial state
  data: function(){
    return {
      searchResults: [],
      query: '',
      navIdx: -1,
    }
  },
  methods: {
    search: function(){
      var self = this;
      let val = self.query.trim()

      self.searchResults.push(val + ' result')

      window.backend.onTextChange(val).then(results => {
        self.searchResults = results
      });
    },

    selectResult: function(resultIdx){
      const clicked = this.searchResults[resultIdx]
      window.backend.Linux.LaunchApplication(clicked.Exec).then(results => {
        console.log(results)
      });
      this.searchResults = []
      this.query = ""
      this.navIdx = -1
    },

    highlightResult: function(resultIdx){
      this.navIdx = resultIdx;
    },

    navigate: function(event){
      switch (event.keyCode) {
        case UP_ARROW_KEY:
          if (this.navIdx >= 0) {
            this.navIdx--;
          }
          break;
        case DOWN_ARROW_KEY:
          if (this.navIdx < this.searchResults.length - 1) {
            this.navIdx++;
          }
          break;
        case ENTER_KEY:
          if (this.navIdx != -1) {
            this.selectResult(this.navIdx);
          }
          break;
      }
    }
  },
}; 
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
[v-cloak] { 
  display: none; 
}

.unselectable {
  user-select: none;
  -moz-user-select: none;
  -khtml-user-select: none;
  -webkit-user-select: none;
  -o-user-select: none;
}

#app {
  font-family: 'Roboto Mono', monospace;
}

#window {
  margin: 0 auto;
  padding: 20px;
  display: block;
  background: #fefefe;
  border-radius: 20px;
}

#gt {
  display: inline-block;
  position: relative;
  font-size: 40px;
  color: #6484b3;
}

#query {
  color: #183053;
  background: none;
  border: none;
  margin: 0;
  display: inline-block;
  font-size: 40px;
  padding: 0 10px 20px 10px;
  -webkit-box-sizing: border-box; 
  -moz-box-sizing: border-box;   
  box-sizing: border-box;   
}

#searchResults {
  margin: 0;
  margin-top: 0;
  font-size: 25px;
  padding-left: 0;
  -webkit-box-sizing: border-box; 
  -moz-box-sizing: border-box;   
  box-sizing: border-box;   
}

.searchResult {
  list-style-type: none;
  list-style-position:inside;
  padding: 20px;
  margin-top: -1px;
  border-style: solid;
  border-color: #6484b3;
  border-width: 1px 0 ;
}

.selectedResult {
  list-style-type: none;
  list-style-position:inside;
  padding: 20px;
  background: #b1c1d9;
  margin-top: -1px;
  border-style: solid;
  border-color: #6484b3;
  border-width: 1px 0 ;
}

.file-icon {
  display: inline-block;
  height: 40px;
}

.file {
  display: inline-block;
}

.filepath {
  font-size: 14px;
  color: #6484b3;
}
</style>
