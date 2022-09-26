FROM ubuntu:latest

LABEL maintainer="nil"

# Run the script sh
# 1. Copy the entrypoint into the image with ADD (add can handle remote url & tar)
ADD docker-entrypoint.sh /

# 2. Change the rights to be able to execute it
RUN chmod u+x /docker-entrypoint.sh

# 3. Run the script to prepare environments
RUN /docker-entrypoint.sh

# Pre-pend the new path
ENV PATH /root/miniconda3/bin:$PATH

# Conda create new environment and install dependencies
# 1. Create & activate environment
RUN conda create -n environment python=3.8
RUN conda activate environment

# 2. Install packages
RUN conda install -y pandas \
    jupyter \
    seaborn \
    statsmodels \
    pystan

# 3. Install fbprophet
RUN conda install -c conda-forge -y fbprophet

# We set which is our workdir and copy the files
WORKDIR /app
COPY * /app/

EXPOSE 8888
RUN conda activate environment

# When run container, at the end CMD execution
# In the case we want an interactive terminal in Python with the environment ocnfigured CMD ["ipython"]
CMD ["jupyter", "notebook", "--port=8888", "--ip=0.0.0.0"]