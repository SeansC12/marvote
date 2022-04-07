import Head from "next/head";
import Layout from "../components/homepage-components/homepage-layout/layout";
import Link from "next/link";
import Sidebar from "../components/sidebar";

export default function Home() {
  return <Sidebar tab="Home" />;
}
