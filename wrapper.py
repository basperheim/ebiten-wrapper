import json
import os
import subprocess

sprite_data = {
  "sprites": [
    {
      "file_path": "assets/forest-tile-128px.png",
      "x": 200,
      "y": 250,
      "animations": [
        # add animations here
      ]
    },
    {
      "file_path": "assets/forest-tile-128px.png",
      "x": 200,
      "y": 250,
      "animations": [
        # add animations here
      ]
    }
  ]
}

with open("sprite.json", "w") as f:
  json.dump(sprite_data, f)

subprocess.run(["go", "run", "main.go", "--sprite", "sprite.json"])