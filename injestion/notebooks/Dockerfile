FROM ubuntu:latest

LABEL maintainer="nil"

# Run the script sh
# 1. Copy the entrypoint into the image with ADD (add can handle remote url & tar)
ADD docker-entrypoint.sh /

# 2. Change the rights to be able to execute it
RUN chmod u+x /docker-entrypoint.sh

# 3. Run the script
RUN /docker-entrypoint.sh

# Pre-pend the nwe path
ENV PATH /root/miniconda3/bin:$PATH

# When run container, at the end CMD execution
CMD ["ipython"]


RUN 