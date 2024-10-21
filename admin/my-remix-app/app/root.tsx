import type { LinksFunction } from "@remix-run/node";
import appStylesHref from "./app.css?url";

export const links: LinksFunction = () => [
  { rel: "stylesheet", href: appStylesHref },
];

import { json } from "@remix-run/node";
import {
  Form,
  Link,
  Links,
  Meta,
  Outlet,
  Scripts,
  NavLink,
  ScrollRestoration,
  useLoaderData,
} from "@remix-run/react";

import { getContacts } from "./data";

export const loader = async () => {
  const contacts = await getContacts();
  return json({ contacts });
};


export default function App() {

  const { contacts } = useLoaderData<typeof loader>();
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
         <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossOrigin="anonymous"/>
      </head>
      <body>


<div className="container-fluid">
    
                <nav className="sidenav">
                  <ul className="nav nav-pills flex-column mb-sm-auto mb-0 align-items-center align-items-sm-start">
                    <li className="nav-item">
                      <Link className="nav-link align-middle px-0" to={`/dashboard`}>Dashboard</Link>
                    </li>
                    <li className="nav-item">
                      <Link className="nav-link align-middle px-0" to={`/posts`}>Posts</Link>
                    </li>
                  </ul>
                </nav>
            
</div>
        <ScrollRestoration />
        <script src="https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js"></script>
      </body>
      
    </html>
  );
}
