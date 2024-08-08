document.addEventListener('DOMContentLoaded', () => {
    const artistsLink = document.getElementById('artists-link');
    const locationsLink = document.getElementById('locations-link');
    const datesLink = document.getElementById('dates-link');
    const relationsLink = document.getElementById('relations-link');
    const content = document.getElementById('content');

    artistsLink.addEventListener('click', (e) => {
        e.preventDefault();
        fetchData('/artists', renderArtists);
    });

    locationsLink.addEventListener('click', (e) => {
        e.preventDefault();
        fetchData('/locations', renderLocations);
    });

    datesLink.addEventListener('click', (e) => {
        e.preventDefault();
        fetchData('/dates', renderDates);
    });

    relationsLink.addEventListener('click', (e) => {
        e.preventDefault();
        fetchData('/relations', renderRelations);
    });

    function fetchData(url, callback) {
        fetch(url)
            .then(response => response.json())
            .then(data => callback(data))
            .catch(error => console.error('Error fetching data:', error));
    }

    function renderArtists(artists) {
        content.innerHTML = '';
        artists.forEach(artist => {
            const card = document.createElement('div');
            card.className = 'card';
            card.innerHTML = `
                <img src="${artist.image}" alt="${artist.name}">
                <h2>${artist.name}</h2>
                <p><strong>Creation Date:</strong> ${artist.creationDate}</p>
                <p><strong>First Album:</strong> ${artist.firstAlbum}</p>
                <p><strong>Members:</strong> ${artist.members.join(', ')}</p>
            `;
            content.appendChild(card);
        });
    }

    function renderLocations(locations) {
        content.innerHTML = '';
        locations.forEach(location => {
            const card = document.createElement('div');
            card.className = 'card';
            card.innerHTML = `
                <h2>Location ID: ${location.id}</h2>
                <p><strong>Locations:</strong> ${location.locations.join(', ')}</p>
            `;
            content.appendChild(card);
        });
    }

    function renderDates(dates) {
        content.innerHTML = '';
        dates.forEach(date => {
            const card = document.createElement('div');
            card.className = 'card';
            card.innerHTML = `
                <h2>Date ID: ${date.id}</h2>
                <p><strong>Dates:</strong> ${date.dates.join(', ')}</p>
            `;
            content.appendChild(card);
        });
    }

    function renderRelations(relations) {
        content.innerHTML = '';
        relations.forEach(relation => {
            const card = document.createElement('div');
            card.className = 'card';
            card.innerHTML = `
                <h2>Relation ID: ${relation.id}</h2>
                <p><strong>Dates and Locations:</strong></p>
                <ul>
                    ${Object.entries(relation.datesLocations).map(([date, locations]) => `
                        <li><strong>${date}:</strong> ${locations.join(', ')}</li>
                    `).join('')}
                </ul>
            `;
            content.appendChild(card);
        });
    }
});
