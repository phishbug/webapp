// index.js
const express = require('express');
const bodyParser = require('body-parser');
const { Client } = require('@elastic/elasticsearch');
const { Client } = require('@opensearch-project/opensearch');
const AWS = require('aws-sdk');
const path = require('path');
var cors = require('cors');


const app = express();
const port = 3000;
app.use(cors());


// Elasticsearch client
const client = new Client({ node: 'http://localhost:9200' });
// Replace with your OpenSearch domain endpoint and region
const openSearchDomain = 'your-opensearch-domain-endpoint'; // e.g., search-your-domain.us-west-1.es.amazonaws.com
const region = 'your-region'; // e.g., 'us-west-1'

// Middleware
app.use(bodyParser.urlencoded({ extended: true }));
app.use(express.static('public'));
app.set('view engine', 'ejs');
app.set('views', path.join(__dirname, 'views'));
app.use(express.json());

// Home route
app.get('/', async (req, res) => {
    const result = await client.search({
        index: 'blogs',
        sort: [{ "_id": { "order": "desc" } }],
        body: {
            query: { match_all: {} }
        }
    });

    res.json(result);
});

// Create document
app.get('/create', (req, res) => {
    res.render('create');
});

app.post('/create', async (req, res) => {
    // res.json(req.body);
    try {
    	result = await client.index({
	        index: 'blogs',
	        body: req.body
		});
        res.json({result});
		console.log(result);
    } catch(e) {
    	res.json({error: e});
    }
    
  
});

// Edit document
app.get('/edit/:id', async (req, res) => {
    const { id } = req.params;
     try {
        const body  = await client.get({
            index: 'blogs',
            id: id,
        });
         res.render('edit', { document: body?._source, id });
    } catch (error) {
        if (error) {
            console.error('Document not found:', error);
        } else {
            console.error('Error retrieving document:', error);
        }
    }

   
});

app.post('/edit/:id', async (req, res) => {
    const { id } = req.params;
    const { title, content } = req.body;
    await client.update({
        index: 'blogs',
        id,
        body: {
            doc: { title, content }
        }
    });
    res.json({okay: 1});
});

// Delete document
app.post('/delete/:id', async (req, res) => {
    const { id } = req.params;
    await client.update({
        index: 'blogs',
        id,
        body: {
            doc: { is_deleted: 1}
        }
    });
    res.json({okay: 1});
});


app.get('/elastic/ping', async (req, res) => {
    res.json({okay: 1});
});

// Start the server
app.listen(port, () => {
    console.log(`Server is running at http://localhost:${port}`);
});
