import requests
import argparse
import os
import random

BASE_URL = 'http://localhost'

def trigger_monitor_action(action):
    api_url = f"{BASE_URL.rstrip('/')}/{action.lstrip('/')}"
    
    try:
        response = requests.get(api_url)
        response.raise_for_status()

        # Generate a random number to append to the filename
        random_number = random.randint(1, 100000)

        # Save the response content to a text file with a random number appended to the filename
        filename = f"response_{action.replace('/', '_')}_{random_number}.txt"
        with open(filename, 'w', encoding='utf-8') as file:
            file.write(response.text)
        print(f"Response saved to file: {filename}")

    except requests.exceptions.RequestException as e:
        print('Error:', str(e))

def main():
    parser = argparse.ArgumentParser(description='Trigger .NET Monitor actions.')
    parser.add_argument('--action', required=True, help='Action to perform in the .NET Monitor')

    args = parser.parse_args()

    trigger_monitor_action(args.action)

if __name__ == '__main__':
    main()
