import Link from "next/link";
import React from "react";
import Sidebar from "../../components/sidebar";

export default function submit_quote() {
  const submitQuote = () => {
    console.log("Submit Quote");
  };

  return <Sidebar tab="Submit" />;
}
