import React from "react";
import Link from "next/link";
import Head from "next/head";

export default function tally() {
  return (
    <div>
      <Head>
        <title>Tally</title>
      </Head>
      <h1>This is the Tally Page</h1>
      <Link href="/">Back To Home</Link>
    </div>
  );
}
