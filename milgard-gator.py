import os
import re
import argparse
from collections import defaultdict
from zeep import Client, Settings
from zeep.transports import Transport
from requests import Session
from requests.auth import HTTPBasicAuth
import urllib3

# Suppress only the single InsecureRequestWarning from urllib3 needed to disable warnings
urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

# Configuration
CUCM_ADDRESS = 'https://10.10.20.1:8443/axl/'
USERNAME = '<redacted>'
PASSWORD = '<redacted>'
WSDL = 'AXLAPI.wsdl'  # Assuming the WSDL file is in the same directory as the script

# Initialize the Zeep client with transport settings
session = Session()
session.verify = False
session.auth = HTTPBasicAuth(USERNAME, PASSWORD)

transport = Transport(session=session, timeout=10)
settings = Settings(strict=False, xml_huge_tree=True)
client = Client(WSDL, settings=settings, transport=transport)

# Function to get all directory numbers
def get_all_dns(debug=False):
    dns = []
    try:
        response = client.service.listLine(searchCriteria={'description': '%'},
                                           returnedTags={'pattern': '', 'description': ''})
        if response['return']:
            dns = response['return']['line']
        if debug:
            print(f"Fetched DNs: {dns}")
    except Exception as e:
        print(f"Error fetching DNs: {e}")
    return dns

# Function to get phones by description
def get_phones_by_description(description_mask, debug=False):
    phones = []
    try:
        response = client.service.listPhone(searchCriteria={'description': f"%{description_mask}%"},
                                            returnedTags={'name': '', 'description': '', 'devicePoolName': ''})
        if response['return']:
            phones = response['return']['phone']
        if debug:
            print(f"Fetched phones: {phones}")
    except Exception as e:
        print(f"Error fetching phones: {e}")
    return phones

# Function to get users by extracted name
def get_users_by_name(extracted_names, debug=False):
    users = []
    for first_name, last_name in extracted_names:
        try:
            response = client.service.listUser(searchCriteria={
                'firstName': f"%{first_name}%",
                'lastName': f"%{last_name}%"
            }, returnedTags={'userid': '', 'firstName': '', 'lastName': ''})
            if response['return']:
                users.extend(response['return']['user'])
            if debug:
                print(f"Fetched users for name {first_name} {last_name}: {response['return']['user']}")
        except Exception as e:
            print(f"Error fetching user for name {first_name} {last_name}: {e}")
    return users

# Function to generate location report
def generate_location_report(debug=False):
    dns = get_all_dns(debug)

    # Group DNs by location
    locations = defaultdict(list)
    for dn in dns:
        description = dn['description']
        if description and isinstance(description, str):
            match = re.match(r'([A-Za-z ]+) - ', description)
            if match:
                location = match.group(1)
                locations[location].append(dn)

    reports = {}

    for location, dns in locations.items():
        phones = get_phones_by_description(location, debug)
        phone_names = [phone['name'] for phone in phones]

        # Extract user names from descriptions
        extracted_names = []
        for dn in dns:
            match = re.search(r' - ([A-Za-z]+) ([A-Za-z]+) - ', dn['description'])
            if match:
                first_name = match.group(1)
                last_name = match.group(2)
                extracted_names.append((first_name, last_name))

        users = get_users_by_name(extracted_names, debug)

        reports[location] = {
            'phones': phones,
            'users': users,
            'directory_numbers': dns
        }

    return reports

# Main function
if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="CUCM AXL Report Generator")
    parser.add_argument('--debug', action='store_true', help='Enable debug mode')
    args = parser.parse_args()

    debug = args.debug
    reports = generate_location_report(debug)

    for location, report in reports.items():
        print(f"\n\033[1;34mReport for Location: {location}\033[0m")
        print("\n\033[1;32mPhones:\033[0m")
        for phone in report['phones']:
            print(f"\033[1;33mName:\033[0m {phone['name']}, \033[1;35mDescription:\033[0m {phone['description']}, \033[1;36mDevice Pool:\033[0m {phone['devicePoolName']}")

        print("\n\033[1;32mUsers:\033[0m")
        for user in report['users']:
            print(f"\033[1;33mUser ID:\033[0m {user['userid']}, \033[1;35mName:\033[0m {user['firstName']} {user['lastName']}")

        print("\n\033[1;32mDirectory Numbers:\033[0m")
        for dn in report['directory_numbers']:
            print(f"\033[1;33mPattern:\033[0m {dn['pattern']}, \033[1;35mDescription:\033[0m {dn['description']}")
