package main

//constants
const start = 97
const end = 122
const base = end - start + 1
const passwordminlength = 5
const passwordmaxlength = 5
const hashToCrack = "$6$hpKW0d9oSGIXHJpn$ItrRNxxgJE.sYixsZtfrDJyYj9XMheFEcyD1ybc/4gJMrICcchyU/D1gYMN7gQKuA3ZDNuqRWbWm37k3zTyvG1"
const MaxNumofRoutines = 50
const RoutineNotify = 2000 // routines log their search at intervals of
//args [0] = minlength , [1] = max length
