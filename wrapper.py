import json
import subprocess

sprite_data = {
    "sprites": [
        {
            "file_path": "assets/forest-tile-128px.png",
            "x": 100,
            "y": 150,
            "animations": [
                # Define animations here if needed
            ]
        },
        {
            "file_path": "assets/forest-tile-128px.png",
            "x": 200,
            "y": 250,
            "animations": [
                # Define animations here if needed
            ]
        }
    ]
}

sprite_json = json.dumps(sprite_data)
subprocess.run(["go", "run", "main.go", "--sprite", sprite_json])
