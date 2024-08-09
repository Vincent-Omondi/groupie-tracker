document.addEventListener("DOMContentLoaded", function() {
    fetchArtists();
});

function fetchArtists() {
    axios.get('/artists')
        .then(response => {
            const artists = response.data;
            displayArtists(artists);
        })
        .catch(error => {
            console.error('Error fetching artists:', error);
        });
}

function displayArtists(artists) {
    const contentGrid = document.getElementById('content-grid');
    contentGrid.innerHTML = '';

    artists.forEach(artist => {
        const card = document.createElement('div');
        card.className = 'content-card';
        card.onclick = () => showArtistDetails(artist.id);

        const img = document.createElement('img');
        img.src = artist.image;
        img.alt = artist.name;
        img.className = 'content-poster';

        const info = document.createElement('div');
        info.className = 'content-info';

        const title = document.createElement('h3');
        title.className = 'content-title';
        title.textContent = artist.name;

        info.appendChild(title);
        card.appendChild(img);
        card.appendChild(info);
        contentGrid.appendChild(card);
    });
}

function showArtistDetails(artistId) {
    axios.get(`/artists/${artistId}`)
        .then(response => {
            const artist = response.data.artist;
            const relation = response.data.relation;

            const modalDetails = document.getElementById('content-modal-details');
            modalDetails.innerHTML = `
                <h1>${artist.name}</h1>
                <img src="${artist.image}" alt="${artist.name}" style="width:100%; height:auto;">
                <p><strong>Creation Date:</strong> ${artist.creationDate}</p>
                <p><strong>First Album:</strong> ${artist.firstAlbum}</p>
                <p><strong>Members:</strong> ${artist.members.join(', ')}</p>
                <h2>Concert Locations</h2>
                <ul>${relation.datesLocations ? Object.keys(relation.datesLocations).map(location => `<li>${location}</li>`).join('') : 'No locations available'}</ul>
                <h2>Concert Dates</h2>
                <ul>${relation.datesLocations ? Object.values(relation.datesLocations).flat().map(date => `<li>${date}</li>`).join('') : 'No dates available'}</ul>
            `;

            document.getElementById('content-modal').style.display = 'block';
        })
        .catch(error => {
            console.error('Error fetching artist details:', error);
        });
}

function closeContentModal() {
    document.getElementById('content-modal').style.display = 'none';
}
