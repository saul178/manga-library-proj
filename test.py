import requests

included_tag_names = ["Action", "Romance"]
excluded_tag_names = ["Harem"]

base_url = "https://api.mangadex.org"

tags = requests.get(f"{base_url}/manga/tag").json()
included_tag_ids = []

# Iterate over each tag in tags["data"]
for tag in tags["data"]:
    # Get the English name of the tag
    tag_name_en = tag["attributes"]["name"]["en"]

    # Check if this tag's name is in the list of included tag names
    if tag_name_en in included_tag_ids:
        # If it is, add its name to the list
        included_tag_ids.append(tag_name_en)


excluded_tag_ids = [
    tag["id"]
    for tag in tags["data"]
    if tag["attributes"]["name"]["en"] in excluded_tag_names
]

print("before tags")
r = requests.get(
    f"{base_url}/manga",
    params={
        "includedTags[]": included_tag_ids,
        "excludedTags[]": excluded_tag_ids,
    },
)
print("after tags")

print([manga["attributes"]["title"] for manga in r.json()["data"]])
