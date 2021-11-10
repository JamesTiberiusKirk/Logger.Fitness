#!/bin/sh
http -v POST :5000/extp jwt-token:$LF_JWT name="Bench press" description="Standart barbel bench press" data_type="sets"
