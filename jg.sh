#!/bin/bash

jg(){
  results=$(justgo "$1")
  echo $results
  cd $(justgo "$results")
}
