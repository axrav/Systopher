import "../styles/globals.css";
import { AppProps } from "next/app";
import { Provider } from "react-redux";
import { wrapper } from "../redux/store";
import Head from "next/head";
import { MantineProvider } from "@mantine/core";

const App = (props: AppProps) => {
  const { Component, pageProps } = props;
  const { store, props: properties } = wrapper.useWrappedStore(wrapper);
  return (
    <>
      <Head>
        <title>Page title</title>
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
      </Head>

      <MantineProvider
        withGlobalStyles
        withNormalizeCSS
        theme={{
          /** Put your mantine theme override here */
          colorScheme: "dark",
        }}
      >
        <Provider store={store}>
          <Component {...properties.pageProps} />
        </Provider>
      </MantineProvider>
    </>
  );
};

export default App;
