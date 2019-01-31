FROM circleci/go:1.11.5 AS bootstrap

COPY GNUMakefile ./

RUN make bootstrap &&\
      rm GNUMakefile

