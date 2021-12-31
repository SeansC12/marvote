import Head from "next/head";
import Image from "next/image";
import styles from "../styles/Home.module.css";
import Layout from "../components/layout/layout";
import Link from "next/link";

export default function Home() {
  return (
    <div>
      <Layout>
        <Head>
          <title>Marvote</title>
        </Head>
        <div>Hello World</div>
      </Layout>
      <Link href="Tally/tally">
        <a>Tally</a>
      </Link>
    </div>
  );
}
