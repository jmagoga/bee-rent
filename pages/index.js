import Head from "next/head";
import styles from "../styles/Home.module.css";
import { ApolloClient, InMemoryCache, gql } from "@apollo/client";

// TODO: fazer formulario, enviar requests. ajeitar informacoes. por propaganda no Google

export default function Home({ bees }) {
  console.log("bees", bees);
  return (
    <div className={styles.container}>
      <Head>
        <title>Polinize</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>
          Compre. Alugue. <a href="https://nextjs.org">Polinize.</a>
        </h1>

        <p className={styles.description}>
          Abelhas selecionadas em Nova Prata, na Serra Gaúcha.
          {/* <code className={styles.code}>pages/index.js</code> */}
        </p>

        <div className={styles.grid}>
          {bees.map((bee) => {
            return (
              <a href={bee.wikiURL} className={styles.card}>
                <h3>{bee.username}</h3>
                <img src={bee.imageURL}></img>
                <p>{bee.info}</p>
              </a>
            );
          })}
        </div>
      </main>

      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Polinize. Nova Prata - RS. copyright
          {/* <img src="/vercel.svg" alt="Vercel Logo" className={styles.logo} /> */}
        </a>
      </footer>
    </div>
  );
}

export async function getStaticProps() {
  const client = new ApolloClient({
    uri: "http://localhost:8081/query",
    cache: new InMemoryCache(),
  });

  const { data } = await client.query({
    query: gql`
      {
        bees {
          id
          username
          info
          wikiURL
          imageURL
        }
      }
    `,
  });

  return {
    props: {
      bees: data.bees,
    },
  };
}
