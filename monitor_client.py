import requests
import argparse
import os

def trigger_monitor_api(base_url, endpoint):
    api_url = f"{base_url.rstrip('/')}/{endpoint.lstrip('/')}"
    
    try:
        response = requests.get(api_url)
        response.raise_for_status()

        content_type = response.headers.get('content-type')

        if content_type and not content_type.startswith('text'):
            # Save the response content to a file
            filename = f"response_{endpoint.replace('/', '_')}"
            with open(filename, 'wb') as file:
                file.write(response.content)
            print(f"Response saved to file: {filename}")
        else:
            # Print the response content
            print('Response:', response.text)

    except requests.exceptions.RequestException as e:
        print('Error:', str(e))

def main():
    parser = argparse.ArgumentParser(description='Trigger .NET Monitor API endpoints.')
    parser.add_argument('--base-url', required=True, help='Base URL of the target .NET application')
    parser.add_argument('--endpoint', required=True, help='Endpoint of the .NET Monitor API to trigger')

    args = parser.parse_args()

    trigger_monitor_api(args.base_url, args.endpoint)

if __name__ == '__main__':
    main()
