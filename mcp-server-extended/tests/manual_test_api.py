#!/usr/bin/env python3
"""Test script to verify Home Assistant API connectivity and automation endpoints.

Run this before setting up the MCP server to ensure everything works.
This is a standalone script, not a pytest test file.
"""

import asyncio
import os
import sys

import aiohttp
from dotenv import load_dotenv

# Only run if executed directly, not when imported by pytest
if __name__ == "__main__":
    load_dotenv()

HA_URL = os.getenv("HA_URL", "http://homeassistant.local:8123")
HA_TOKEN = os.getenv("HA_TOKEN", "")

if __name__ == "__main__":
    if not HA_TOKEN:
        print("ERROR: HA_TOKEN not set. Please set it in .env or environment variables.")
        sys.exit(1)


async def test_connection():
    """Test basic connection to Home Assistant."""
    # Get values at runtime to avoid import-time errors
    load_dotenv()
    ha_url = os.getenv("HA_URL", "http://homeassistant.local:8123")
    ha_token = os.getenv("HA_TOKEN", "")

    print(f"Testing connection to {ha_url}...")

    try:
        async with aiohttp.ClientSession() as session:
            headers = {"Authorization": f"Bearer {ha_token}"}

            # Test API root
            async with session.get(f"{ha_url}/api/", headers=headers) as response:
                if response.status == 200:
                    print("✓ API connection successful")
                    data = await response.json()
                    print(f"  Home Assistant version: {data.get('message', 'unknown')}")
                else:
                    print(f"✗ API connection failed: {response.status}")
                    return False

            # Test automation endpoint
            async with session.get(f"{ha_url}/api/automation", headers=headers) as response:
                if response.status == 200:
                    automations = await response.json()
                    print(f"✓ Automation API accessible")
                    print(f"  Found {len(automations)} automations")

                    if automations:
                        print("\n  Sample automations:")
                        for auto in automations[:3]:
                            print(
                                f"    - {auto.get('alias', auto.get('id', 'Unknown'))} "
                                f"(ID: {auto.get('id', 'N/A')}, "
                                f"Enabled: {auto.get('enabled', True)})"
                            )
                    return True
                else:
                    print(f"✗ Automation API failed: {response.status}")
                    text = await response.text()
                    print(f"  Response: {text}")
                    return False

    except aiohttp.ClientError as e:
        print(f"✗ Connection error: {e}")
        print("\nTroubleshooting:")
        print("  1. Check HA_URL is correct")
        print("  2. Verify Home Assistant is running")
        print("  3. Check network connectivity")
        return False
    except Exception as e:
        print(f"✗ Unexpected error: {e}")
        return False


async def test_create_automation():
    """Test creating a test automation."""
    load_dotenv()
    ha_url = os.getenv("HA_URL", "http://homeassistant.local:8123")
    ha_token = os.getenv("HA_TOKEN", "")

    print("\nTesting automation creation...")

    test_automation = {
        "alias": "MCP Test Automation",
        "description": "Test automation created by MCP server",
        "trigger": [
            {
                "platform": "time",
                "at": "12:00:00",
            }
        ],
        "action": [
            {
                "service": "system_log.write",
                "data": {
                    "message": "MCP test automation triggered",
                },
            }
        ],
        "mode": "single",
    }

    try:
        async with aiohttp.ClientSession() as session:
            headers = {
                "Authorization": f"Bearer {ha_token}",
                "Content-Type": "application/json",
            }

            async with session.post(
                f"{ha_url}/api/automation",
                headers=headers,
                json=test_automation,
            ) as response:
                if response.status in [200, 201]:
                    result = await response.json()
                    print(f"✓ Test automation created successfully")
                    print(f"  Automation ID: {result.get('id', 'N/A')}")
                    return result.get("id")
                else:
                    text = await response.text()
                    print(f"✗ Failed to create automation: {response.status}")
                    print(f"  Response: {text}")
                    return None
    except Exception as e:
        print(f"✗ Error creating automation: {e}")
        return None


async def test_delete_automation(automation_id):
    """Test deleting the test automation."""
    load_dotenv()
    ha_url = os.getenv("HA_URL", "http://homeassistant.local:8123")
    ha_token = os.getenv("HA_TOKEN", "")

    if not automation_id:
        return

    print(f"\nCleaning up test automation ({automation_id})...")

    try:
        async with aiohttp.ClientSession() as session:
            headers = {"Authorization": f"Bearer {ha_token}"}

            async with session.delete(
                f"{ha_url}/api/automation/{automation_id}",
                headers=headers,
            ) as response:
                if response.status == 200:
                    print("✓ Test automation deleted successfully")
                else:
                    print(f"⚠ Could not delete test automation: {response.status}")
    except Exception as e:
        print(f"⚠ Error deleting automation: {e}")


async def main():
    """Run all tests."""
    print("=" * 60)
    print("Home Assistant Automation API Test")
    print("=" * 60)
    print()

    # Test connection
    if not await test_connection():
        print("\n❌ Connection test failed. Please fix issues before proceeding.")
        sys.exit(1)

    # Test creating automation (optional)
    print("\n" + "=" * 60)
    create_test = input("Create a test automation? (y/N): ").strip().lower()
    if create_test == "y":
        automation_id = await test_create_automation()
        if automation_id:
            delete_test = input("\nDelete test automation? (Y/n): ").strip().lower()
            if delete_test != "n":
                await test_delete_automation(automation_id)

    print("\n" + "=" * 60)
    print("✓ All tests passed! You can now use the MCP server.")
    print("=" * 60)


if __name__ == "__main__":
    asyncio.run(main())
