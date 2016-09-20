{{define "title"}}Pokemon Name Generator{{end}}
{{define "content"}}
<div class="row">
    <div class="generator">
        <div class="inner">
            <h2>Your Pokemon's Name Is</h2>
            <div id="generatedName"><span></span></div> 
            <div class="name-actions">
                <button id="generateName">New Name</button>
                <button id="addToFavorites">&#9733;</button>
            </div>
        </div>
    </div>
    <div class="favorites">
        <div class="inner">
            <h2>Favorites</h2>
            <ul id="favoritesList"></ul>
        </div>
    </div>
</div>
{{end}}