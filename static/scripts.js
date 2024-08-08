const API_KEY = '4ea270f32fe4e8fcdfd68b4cd5a7074f';
const BASE_URL = 'https://api.themoviedb.org/3';
const IMG_BASE_URL = 'https://image.tmdb.org/t/p/w500';
let currentType = 'movie';

async function fetchContent() {
    try {
        const url = `${BASE_URL}/trending/${currentType}/week?api_key=${API_KEY}&language=en-US`;
        const response = await axios.get(url);
        const content = response.data.results;
        updateUI(content);
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}

function updateUI(content) {
    const contentGrid = document.getElementById('content-grid');
    contentGrid.innerHTML = '';

    content.forEach(item => {
        const contentCard = document.createElement('div');
        contentCard.className = 'content-card';
        contentCard.innerHTML = `
            <img class="content-poster" src="${IMG_BASE_URL}${item.poster_path}" alt="${item.title || item.name}">
            <div class="content-info">
                <h2 class="content-title">${item.title || item.name}</h2>
                <p class="content-details">Rating: ${item.vote_average}/10</p>
            </div>
        `;
        contentCard.addEventListener('click', () => openContentModal(item.id));
        contentGrid.appendChild(contentCard);
    });
}

async function openContentModal(contentId) {
    try {
        const response = await axios.get(`${BASE_URL}/${currentType}/${contentId}?api_key=${API_KEY}&language=en-US&append_to_response=credits,similar,videos`);
        const content = response.data;

        let trailerKey = '';
        if (content.videos && content.videos.results.length > 0) {
            const trailer = content.videos.results.find(video => video.type === 'Trailer') || content.videos.results[0];
            trailerKey = trailer.key;
        }

        const modalContent = document.getElementById('content-modal-details');
        modalContent.innerHTML = `
            <h2>${content.title || content.name}</h2>
            ${trailerKey ? `
                <div class="video-container">
                    <iframe src="https://www.youtube.com/embed/${trailerKey}" frameborder="0" allow="autoplay; encrypted-media" allowfullscreen></iframe>
                </div>
            ` : `
                <img src="${IMG_BASE_URL}${content.poster_path}" alt="${content.title || content.name}" style="width: 200px;">
            `}
            <p>${content.overview}</p>
            <p>Release date: ${content.release_date || content.first_air_date}</p>
            <p>Rating: ${content.vote_average}/10</p>
            <h3>Cast:</h3>
            <p>${content.credits.cast.slice(0, 5).map(actor => actor.name).join(', ')}</p>
            <h3>Similar content:</h3>
            <div style="display: flex; overflow-x: auto;">
                ${content.similar.results.slice(0, 5).map(item => `
                    <div style="margin-right: 10px;">
                        <img src="${IMG_BASE_URL}${item.poster_path}" alt="${item.title || item.name}" style="width: 100px;">
                        <p>${item.title || item.name}</p>
                    </div>
                `).join('')}
            </div>
        `;

        document.getElementById('content-modal').style.display = 'block';
    } catch (error) {
        console.error('Error fetching content details:', error);
    }
}

function closeContentModal() {
    document.getElementById('content-modal').style.display = 'none';
}

async function searchContent() {
    const query = document.getElementById('search-input').value;
    if (query.trim() === '') return;

    try {
        const response = await axios.get(`${BASE_URL}/search/${currentType}?api_key=${API_KEY}&language=en-US&query=${encodeURIComponent(query)}`);
        const content = response.data.results;
        updateUI(content);
    } catch (error) {
        console.error('Error searching content:', error);
    }
}

function switchTab(type) {
    currentType = type === 'movies' ? 'movie' : 'tv';
    document.querySelectorAll('.tab').forEach(tab => tab.classList.remove('active'));
    event.target.classList.add('active');
    fetchContent();
}

document.getElementById('search-input').addEventListener('keypress', (event) => {
    if (event.key === 'Enter') {
        searchContent();
    }
});

window.onclick = function(event) {
    if (event.target == document.getElementById('content-modal')) {
        closeContentModal();
    }
}

fetchContent();
