import React from 'react';

const CodePreview = () => {
    const codeSnippet = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Document</title>
</head>
<body>
    <h1>Hello, World!</h1>
</body>
</html>`;

    const copyToClipboard = () => {
        navigator.clipboard.writeText(codeSnippet).then(() => {
            alert('Code copied to clipboard!');
        }).catch(err => {
            console.error('Could not copy code: ', err);
        });
    };

    return (
        <div className="container mt-5">
            <div className="card">
                <div className="card-header">
                    Code Preview
                </div>
                <div className="card-body">
                    <pre className="bg-light border p-3">
                        <code>
                            {codeSnippet}
                        </code>
                    </pre>
                    <button className="btn btn-primary mt-3" onClick={copyToClipboard}>
                        Copy Code
                    </button>
                </div>
            </div>
        </div>
    );
};

export default CodePreview;
