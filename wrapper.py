import json
import subprocess

sprite_data = {
    "sprites": [
        {
            "file_path": "assets/forest-tile-128px.png",
            "x": 100,
            "y": 150,
            "animations": [
                {
                    "type": "move",
                    "start_x": 100,
                    "end_x": 300,
                    "speed": 2
                }
            ]
        },
        {
            "file_path": "assets/forest-tile-128px.png",
            "x": 200,
            "y": 250,
            "animations": [
                {
                    "type": "move",
                    "start_x": 200,
                    "end_x": 400,
                    "speed": 3
                }
            ]
        }
    ]
}

sprite_json = json.dumps(sprite_data)
subprocess.run(["go", "run", "main.go", "--sprite", sprite_json])
